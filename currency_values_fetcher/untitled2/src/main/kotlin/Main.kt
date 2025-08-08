package org.example

import org.example.domains.task_executor.Task
import org.example.domains.task_executor.TaskExecutor
import java.time.Duration

fun main() {
    val fetchCurrenciesTask = Task {
    }
    val fetchCurrenciesInterval = Duration.ofSeconds(32)
    val taskExecutor = TaskExecutor(fetchCurrenciesTask, fetchCurrenciesInterval)

    taskExecutor.execute()
}