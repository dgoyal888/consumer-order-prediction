package util

import "github.com/google/uuid"


func GenerateUUID() (string,error) {
  u2,err  := uuid.NewUUID()

  if err != nil {
    return "",err
  }

  return u2.String(),nil
}