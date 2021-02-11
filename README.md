# teamlab_task
front-end: ios app. using swift, swiftUI, URLSession.
backend: using go, gin.

How to use:
1, front-end: open "ios front-end" on xcode, click "build and then run the current scheme".
2, back-end: "go run backend.go" on terminal.
3, front-end: input the id of a teacher (like "tc1") or a student (like "st001"), click the button "発行" on the iphone screen to create a verticifation code. The code will be displayed on terminal.
4, front-end: input the code, click the button "ログイン".
5, front-end: On the right page "個人情報", you can view the personal information of the one who owns the id. On the left page "成績", if a student owns the id, you can view the scores of this student. If a teacher owns the id, you can view the students' list, and check the scores of each student by click on the row where the student is.
