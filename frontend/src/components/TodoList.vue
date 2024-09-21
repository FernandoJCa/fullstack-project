<template>
  <div class="max-w-md mx-auto mt-10 p-6 bg-white dark:bg-gray-800 rounded-lg shadow-md transition-colors duration-200">
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-3xl font-bold text-gray-900 dark:text-white">Todo List</h1>
      <button @click="toggleDarkMode" class="p-2 rounded-full bg-gray-200 dark:bg-gray-600 transition-colors duration-200">
        <svg v-if="isDarkMode" xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-yellow-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
        </svg>
        <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-gray-800" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z" />
        </svg>
      </button>
    </div>

    <!-- Feedback messages -->
    <div v-if="feedback.message" :class="['mb-4 p-2 rounded', feedback.type === 'success' ? 'bg-green-100 text-green-700 dark:bg-green-700 dark:text-green-100' : 'bg-red-100 text-red-700 dark:bg-red-700 dark:text-red-100']">
      {{ feedback.message }}
    </div>

    <div class="mb-4 flex">
      <input
          v-model="newTodo"
          @keyup.enter="addTodo"
          placeholder="Add a new todo"
          class="flex-grow p-2 border border-gray-300 dark:border-gray-600 rounded-l bg-white dark:bg-gray-700 text-gray-900 dark:text-white transition-colors duration-200"
      >
      <button @click="addTodo" class="px-4 py-2 bg-green-500 text-white rounded-r hover:bg-green-600 transition-colors duration-200">
        Add
      </button>
    </div>
    <div class="bg-white dark:bg-gray-700 shadow-md rounded transition-colors duration-200">
      <TodoItem
          v-for="todo in todos"
          :key="todo.id"
          :todo="todo"
          @toggle-complete="toggleComplete"
          @edit-todo="startEdit"
          @delete-todo="deleteTodo"
      />
    </div>
    <div v-if="editingTodo" class="mt-4 p-4 bg-gray-100 dark:bg-gray-600 rounded transition-colors duration-200">
      <input
          v-model="editingTodo.title"
          @keyup.enter="updateTodo"
          class="w-full p-2 mb-2 border border-gray-300 dark:border-gray-500 rounded bg-white dark:bg-gray-700 text-gray-900 dark:text-white transition-colors duration-200"
      >
      <div class="flex justify-end">
        <button @click="cancelEdit" class="px-4 py-2 mr-2 bg-gray-300 dark:bg-gray-500 text-gray-800 dark:text-white rounded hover:bg-gray-400 dark:hover:bg-gray-400 transition-colors duration-200">Cancel</button>
        <button @click="updateTodo" class="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 transition-colors duration-200">Update</button>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted, watch } from 'vue'
import axios from 'axios'
import TodoItem from './TodoItem.vue'

export default {
  components: { TodoItem },
  setup() {
    const todos = ref([])
    const newTodo = ref('')
    const editingTodo = ref(null)
    const isDarkMode = ref(localStorage.getItem('darkMode') === 'true')
    const feedback = ref({ message: '', type: 'success' })

    const showFeedback = (message, type = 'success') => {
      feedback.value = { message, type }
      setTimeout(() => {
        feedback.value = { message: '', type: 'success' }
      }, 3000)
    }

    const fetchTodos = async () => {
      try {
        const response = await axios.get('http://localhost:8080/todos')
        todos.value = response.data.data
        showFeedback(`${todos.value.length} todos loaded successfully`)
      } catch (error) {
        console.error('Error fetching todos:', error)
        showFeedback('Failed to load todos', 'error')
      }
    }

    const addTodo = async () => {
      if (newTodo.value.trim()) {
        try {
          const response = await axios.post('http://localhost:8080/todos', {
            title: newTodo.value,
            completed: false
          })
          todos.value.push(response.data.data)
          newTodo.value = ''
          showFeedback('Todo created successfully')
        } catch (error) {
          console.error('Error adding todo:', error)
          showFeedback('Failed to create todo', 'error')
        }
      }
    }

    const toggleComplete = async (id) => {
      const todo = todos.value.find(t => t.id === id)
      if (todo) {
        try {
          await axios.put(`http://localhost:8080/todos/${id}`, {
            ...todo,
            completed: !todo.completed
          })
          todo.completed = !todo.completed
          showFeedback(`Todo ${todo.completed ? 'completed' : 'uncompleted'}`)
        } catch (error) {
          console.error('Error updating todo:', error)
          showFeedback('Failed to update todo', 'error')
        }
      }
    }

    const startEdit = (todo) => {
      editingTodo.value = {...todo}
    }

    const cancelEdit = () => {
      editingTodo.value = null
    }

    const updateTodo = async () => {
      if (editingTodo.value) {
        try {
          await axios.put(`http://localhost:8080/todos/${editingTodo.value.id}`, editingTodo.value)
          const index = todos.value.findIndex(t => t.id === editingTodo.value.id)
          if (index !== -1) {
            todos.value[index] = {...editingTodo.value}
          }
          editingTodo.value = null
          showFeedback('Todo updated successfully')
        } catch (error) {
          console.error('Error updating todo:', error)
          showFeedback('Failed to update todo', 'error')
        }
      }
    }

    const deleteTodo = async (id) => {
      try {
        await axios.delete(`http://localhost:8080/todos/${id}`)
        todos.value = todos.value.filter(t => t.id !== id)
        showFeedback('Todo deleted successfully')
      } catch (error) {
        console.error('Error deleting todo:', error)
        showFeedback('Failed to delete todo', 'error')
      }
    }

    const toggleDarkMode = () => {
      isDarkMode.value = !isDarkMode.value
      localStorage.setItem('darkMode', isDarkMode.value)
    }

    watch(isDarkMode, (newValue) => {
      if (newValue) {
        document.documentElement.classList.add('dark')
      } else {
        document.documentElement.classList.remove('dark')
      }
    }, {immediate: true})

    onMounted(fetchTodos)

    return {
      todos,
      newTodo,
      editingTodo,
      isDarkMode,
      feedback,
      addTodo,
      toggleComplete,
      startEdit,
      cancelEdit,
      updateTodo,
      deleteTodo,
      toggleDarkMode
    }
  }
}
</script>