package mysql

import (
	"database/sql"
	"github.com/Suellen-Kitten/PF_CC5M-WEB/pkg/models"
)

type SnippetModel struct{
  DB *sql.DB
}

func(m *SnippetModel)Insert(title, content, expired string) (int, error){
  stmt := `INSERT INTO Snippets (title, conten, created, expires)
            VALUES(?,?,UTC_TIMESTAMP(), DATA_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

  result, err := m.DB.Exec(stmt, title, content, expired)
  if err != nil{
    return 0, err
  }
  id, err := result.LastInsertId()
  if err != nil{
    return 0, err
  }
  return int(id),nil
}

func(m *SnippetModel) Get(id int)(*models.Snippet, error){
  return nil, nil
}

func(m *SnippetModel) Latest()([]*models.Snippet, error){
  return nil, nil
}