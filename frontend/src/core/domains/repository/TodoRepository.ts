import type { Todo } from '../models/todo';

const baseURL = 'http://localhost:8080';

export class TodoRepository {
	static async getAll(): Promise<Todo[]> {
		const response = await fetch(baseURL + '/todos');
		const todos: Todo[] = await response.json();
		return todos || [];
	}

	static async addTodo(todo: Todo): Promise<void> {
		await fetch(baseURL + '/todo/store', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify(todo)
		});
	}

	static async deleteTodo(id: Pick<Todo, 'id'>): Promise<void> {
		await fetch(baseURL + '/todo/delete', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ id: id })
		});
	}
}
