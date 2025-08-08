package org.example.domains.task_executor

class Task(
    private val task: () -> Unit
) {
    fun execute() {
        task()
    }
}