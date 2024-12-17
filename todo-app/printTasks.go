package printTasks
func printTasks(taskItems []string) {
	fmt.Println("List of Todos");
	for index, task := range taskItems {
		fmt.Printf("%d: %s\n", index+1, task)
	}
}