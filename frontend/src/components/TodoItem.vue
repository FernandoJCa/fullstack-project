<template>
  <div class="flex items-center justify-between p-4 border-b border-gray-200 dark:border-gray-700 transition-colors duration-200">
    <div class="flex items-center flex-grow">
      <div class="relative inline-block w-10 mr-2 align-middle select-none transition duration-200 ease-in">
        <input
            type="checkbox"
            :id="'toggle-' + todo.id"
            :checked="todo.completed"
            @change="$emit('toggle-complete', todo.id)"
            class="toggle-checkbox absolute block w-6 h-6 rounded-full bg-white border-4 appearance-none cursor-pointer"
        />
        <label
            :for="'toggle-' + todo.id"
            class="toggle-label block overflow-hidden h-6 rounded-full bg-gray-300 cursor-pointer"
        ></label>
      </div>
      <span :class="{ 'line-through text-gray-500 dark:text-gray-400': todo.completed, 'text-gray-800 dark:text-gray-200': !todo.completed }">
        {{ todo.title }}
      </span>
    </div>
    <div class="flex items-center">
      <button @click="$emit('edit-todo', todo)" class="px-2 py-1 bg-blue-500 text-white rounded hover:bg-blue-600 mr-2 transition duration-150 ease-in-out">
        Edit
      </button>
      <button @click="$emit('delete-todo', todo.id)" class="px-2 py-1 bg-red-500 text-white rounded hover:bg-red-600 transition duration-150 ease-in-out">
        Delete
      </button>
    </div>
  </div>
</template>

<script>
export default {
  props: ['todo'],
  emits: ['toggle-complete', 'edit-todo', 'delete-todo']
}
</script>

<style scoped>
.toggle-checkbox:checked {
  right: 0;
  border-color: #68D391;
}

.toggle-checkbox:checked + .toggle-label {
  background-color: #68D391;
}
</style>