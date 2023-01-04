<script context="module" lang="ts">
	import type { Load } from '@sveltejs/kit';
	import { TodoRepository } from '../core/domains/repository/TodoRepository';

	export const load: Load = async () => {

	}
</script>

<script lang="ts">
	import TodoCard from '../component/TodoCard/index.svelte'
	import { Input, Button } from 'flowbite-svelte'
</script>

<section>
	<h2 class="mb-8 font-bold text-4xl">TODO</h2>
	<form class="w-full mb-8">
		<Input placeholder="テキストを入力してください" size="lg">
			<Button slot="right" size="sm" type="submit">追加</Button>
		</Input>
	</form>
	<TodoCard />
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
