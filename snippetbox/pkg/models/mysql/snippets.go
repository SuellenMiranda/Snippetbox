package mysql

import (
	"database/sql"
  	"github.com/Suellen-Kitten/PF_CC5M-WEB/pkg/models"
)

type SnippetModel struct{
  DB *sql.DB
}

func(m *SnippetModel)Insert(title, content, expired string) (int, error){
  stmt := `INSERT INTO Snippets (title, content, created, expires)
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
  stmt := `SELECT id, title, content, create, expires FROM snippets
           WHERE expires > UTC_TIMESTAMP() AND id = ?`
  row := m.DB.QueryRow(stmt, id)

  s := &models.Snippet{}

  err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Create, &s.Expires)
  
  if err == sql.ErrNoRows{
    return nil, models.ErrNoRecord
  }else if err != nil{
    return nil, err
  }
    
  return s, nil
}

func(m *SnippetModel) Latest()([]*models.Snippet, error){
  stmt := `SELECT id, title, content, create, expires FROM snippets   
           WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

  rows, err := m.DB.Query(stmt)
  if err != nil{
    return nil, err
  }
  defer rows.Close()

  snippets :=  []*models.Snippet{}
  for rows.Next(){
    s := &models.Snippet{}
    err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Create, &s.Expires)
    if err != nil{
      return nil, err
    }
    snippets = append(snippets, s)
  }
  err = rows.Err()
  if err != nil{
    return nil, err
  }
  
  return snippets, nil
}