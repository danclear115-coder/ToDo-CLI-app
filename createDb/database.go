package database

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var ErrNoId = errors.New("Задача с таким Id не найдена")

var DB *gorm.DB

func databasePath() string {
	_, thisFile, _, _ := runtime.Caller(0)

	projectRoot := filepath.Dir(filepath.Dir(thisFile))
	return filepath.Join(projectRoot, "database", "database.db")
}

type Task struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Title       string `gorm:"not null"`
	Content     string
	Priority    string
	IsCompleted bool `gorm:"not null;default:false"`
}

func InitDB() {
	dbPath := databasePath()
	databaseDir := filepath.Dir(dbPath)

	if err := os.MkdirAll(databaseDir, 0755); err != nil {
		log.Fatal("create database directory:", err)
	}

	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	err = DB.AutoMigrate(&Task{})
	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	log.Println("Database is ready with GORM")
}

func IdCheck(id uint) error {

	var currentUser Task

	err := DB.Where("id = ?", id).First(&currentUser).Error

	if err != nil {
		return ErrNoId
	}

	return err
}

func CreateTask(title, content, priority string) error {
	newTask := Task{
		Title:       title,
		Content:     content,
		Priority:    priority,
		IsCompleted: false,
	}

	result := DB.Create(&newTask)
	return result.Error
}

func GetAllTasks() ([]Task, error) {
	var tasks []Task
	result := DB.Find(&tasks)
	return tasks, result.Error
}

func GetUncompletedTasks() ([]Task, error) {

	var uncompleteTasks []Task
	result := DB.Where("is_completed = ?", false).Find(&uncompleteTasks)
	return uncompleteTasks, result.Error

}

func CompleteTask(id uint) error {

	if err := IdCheck(id); err != nil {
		return err
	}

	var currentTaskIsCompleted bool
	err := DB.Model(&Task{}).Where("id = ?", id).Select("IsCompleted").Scan(&currentTaskIsCompleted).Error

	if err != nil {
		fmt.Println("Task don't find or error")
	}

	result := DB.Model(&Task{}).Where("id = ?", id).Update("IsCompleted", !currentTaskIsCompleted)
	return result.Error
}

func ChangeTask(title string, content string, priority string, id uint) error {

	if err := IdCheck(id); err != nil {
		return err
	}

	result := DB.Model(&Task{}).Where("id = ?", id).Updates(map[string]interface{}{
		"title":    title,
		"content":  content,
		"priority": priority,
	})

	return result.Error

}

func DeleteTask(id uint) error {
	result := DB.Delete(&Task{}, id)
	return result.Error
}
