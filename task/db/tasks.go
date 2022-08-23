package db

import (
	"encoding/binary"
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var completedBucket = []byte("completed")
var db *bolt.DB

type Task struct {
	Key       uint64
	Value     string
	Created   time.Time
	Completed time.Time
}

func Init(dbPath string) (err error) {
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		_, err = tx.CreateBucketIfNotExists(completedBucket)
		return err
	})
}

func CreateTask(task *Task) (err error) {
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		key := itob(id64)
		encoded, err := json.Marshal(task)
		if err != nil {
			return err
		}
		return b.Put(key, []byte(encoded))
	})
	if err != nil {
		return err
	}
	return nil
}

func CompleteTask(task *Task) (err error) {
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(completedBucket)
		id64, _ := b.NextSequence()
		key := itob(id64)
		encoded, err := json.Marshal(task)
		if err != nil {
			return err
		}
		return b.Put(key, []byte(encoded))
	})
	if err != nil {
		return err
	}
	return nil
}

func ReadTasks() (tasks []*Task, err error) {
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			task := &Task{}
			err := json.Unmarshal(v, task)
			if err != nil {
				return err
			}
			tasks = append(tasks, &Task{
				Key:     btoi(k),
				Value:   task.Value,
				Created: task.Created,
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func ReadCompleted() (tasks []*Task, err error) {
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(completedBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			task := &Task{}
			err := json.Unmarshal(v, task)
			if err != nil {
				return err
			}
			tasks = append(tasks, &Task{
				Key:       btoi(k),
				Value:     string(task.Value),
				Created:   task.Created,
				Completed: task.Completed,
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func ReadCompletedLastMinute() (tasks []*Task, err error) {
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(completedBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			task := &Task{}
			err := json.Unmarshal(v, task)
			if err != nil {
				return err
			}
			if task.Completed.Before(time.Now().Add(-5 * time.Minute)) {
				continue
			}
			tasks = append(tasks, &Task{
				Key:       btoi(k),
				Value:     string(task.Value),
				Created:   task.Created,
				Completed: task.Completed,
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func DeleteTask(key uint64) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}

func itob(val uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, val)
	return b
}

func btoi(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}
