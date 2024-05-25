import { defineStore } from 'pinia'


export default todos = defineStore('todos', {
    state: () => ({
      todos: [],
      filter: 'all',
      nextId: 0,
    }),
    actions: {
    }
})



