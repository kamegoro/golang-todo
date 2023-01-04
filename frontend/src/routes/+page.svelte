<script lang="ts">
	import { TodoRepository } from '../core/domains/repository/TodoRepository';
	import { Input, Button, Spinner, Alert, Tabs, TabItem, Checkbox } from 'flowbite-svelte';
	import { onMount } from 'svelte';
	import type { Todo } from 'src/core/domains/models/todo';

	let todos: Todo[] | null = null;
	let text: string = '';
	let isLoading: boolean = false;

	let unCompletedTodos: Todo[] = [];
	let completedTodos: Todo[] = [];

	onMount(async () => {
		todos = await TodoRepository.getAll();
		unCompletedTodos = todos.filter(v => !v.completed);
		completedTodos = todos.filter(v => v.completed)
	});

	const addTodo = async () => {
		isLoading = true;
		await TodoRepository.addTodo({
			id: Date.now(),
			text,
			completed: false
		})

		todos = await TodoRepository.getAll()
		text = "";
		isLoading = false
	};

</script>

<section>
	<h2 class="mb-8 font-bold text-4xl">TODO</h2>
	<form class="w-full mb-2">
		<Input placeholder="テキストを入力してください" size="lg" bind:value={text}>
			<Button slot="right" size="sm" type="submit" disabled={!text} on:click={async () => await addTodo()}>
				{#if isLoading}
					<Spinner class="mr-3" size="4" color="white" />
				{/if}
				追加
			</Button>
		</Input>
	</form>
	{#if !todos}
		<Spinner size={'14'} />
	{:else if !todos.length}
		<Alert>
			<span slot="icon"
				><svg
					aria-hidden="true"
					class="w-5 h-5"
					fill="currentColor"
					viewBox="0 0 20 20"
					xmlns="http://www.w3.org/2000/svg"
					><path
						fill-rule="evenodd"
						d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"
						clip-rule="evenodd"
					/></svg
				>
			</span>
			<span class="font-medium">TODOリストは0件です。</span>
		</Alert>
	{:else}
		<Tabs style="underline" class="mb-6 w-96">
			<TabItem open title="TODOリスト">
				{#each unCompletedTodos as todo}
					<div class="flex items-center w-full px-6 border-r-8 bg-slate-200 p-2 mb-2">
						<Checkbox checked={todo.completed}/>
						<p class="ml-2">{todo.text}</p>
					</div>
				{/each}
			</TabItem>
			<TabItem title="完了リスト">
				{#each completedTodos as todo}
					<div class="flex items-center w-full px-6 border-r-2 bg-slate-400 p-2 mb-2">
						<Checkbox checked={false} disabled />
						<p class="ml-2">{todo.text}</p>
					</div>
				{/each}
			</TabItem>
		</Tabs>
	{/if}
</section>

<style>
	section {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		flex: 0.6;
	}
</style>
