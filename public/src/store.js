import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex);

const store = new Vuex.Store({
    state: {
        tasks: [
            {
                id: "1",
                name: 'Task 1',
                description: 'Exponential number'
            },
            {
                id: "2",
                name: 'Task 2',
                description: 'Lexer'
            },
            {
                id: "3",
                name: 'Task 3',
                description: 'Parser'
            }
        ]
    },
    actions: {

    },
    mutations: {
        addTask (state, task) {
            state.tasks.push(task)
        }
    },
    getters: {
            tasks: state => {
                return state.tasks
            },
            getTaskById: (state) => (id) => {
                return state.tasks.find(x => x.id === id)
            },
            taskCount: state => {
                return state.tasks.length
            }
    }
});

export default store
