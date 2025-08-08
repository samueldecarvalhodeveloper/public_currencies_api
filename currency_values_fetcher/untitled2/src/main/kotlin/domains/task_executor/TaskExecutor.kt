package org.example.domains.task_executor

import java.time.Duration

class TaskExecutor(private val taskToExecute: Task, private val intervals: Duration) {
    fun execute() {
        while (true) {
            taskToExecute.execute()

            Thread.sleep(intervals)
        }
    }
}