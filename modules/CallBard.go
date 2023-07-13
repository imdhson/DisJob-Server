package modules

import (
	"bufio"
	"context"
	"os"
	"os/exec"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CallBard(oid primitive.ObjectID, title string, content string) {
	//만들었던 파일 삭제
	err := os.Remove("CallBard/input.txt")
	ErrOK(err)
	err = os.Remove("CallBard/output.txt")
	ErrOK(err)

	//AI에게 페르소나 부여
	input_refined := `제목 : ` + title + `내용 : ` + content

	//파일에다가 인풋 작성
	text_file, err := os.OpenFile("CallBard/input.txt", os.O_CREATE|os.O_WRONLY, 0644)
	ErrOK(err)
	defer text_file.Close()

	_, err = text_file.WriteString(input_refined)
	ErrOK(err)

	//파이썬 CallBard.py 호출
	cmd := exec.Command("python", "CallBard/CallBard.py")
	cmd.Run()

	time.Sleep(5 * time.Second) //5초간 대기

	//output.txt 읽어들이기
	output_file, err := os.Open("CallBard/output.txt")
	ErrOK(err)
	defer output_file.Close()

	var ai_content string
	output_file_scanner := bufio.NewScanner(output_file)
	for output_file_scanner.Scan() {
		ai_content += output_file_scanner.Text()
	}

	//호출 완료
	godotenv.Load()
	URI := os.Getenv("MONGODB_URI")
	if URI == "" {
		Critical(err)
	}
	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	ErrOK(err)
	defer func() {
		err := db.Disconnect(context.TODO())
		Critical(err)
	}()

	//댓글 등록
	coll := db.Database("dj_board").Collection("comments")
	comments_struct := Dj_board_comments{
		Djjobsid: oid,
		Djuserid: primitive.NilObjectID,
		CreateAt: time.Now(),
		Content:  ai_content,
		GenbyAI:  true,
	}
	_, err = coll.InsertOne(context.TODO(), comments_struct)
	ErrOK(err)
}
