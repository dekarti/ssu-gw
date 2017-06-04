import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex);

const store = new Vuex.Store({
    state: {
        tasks: []
    },
    actions: {
        fetchTasks ({ commit }) {
            axios.get("/tasks")
                .then((response) => {
                    commit('setTasks', response.data.tasks)
                })
        },
        addTask ({ commit }, task) {
            axios.post("/tasks", task)
                .then((response) => {
                    commit('addTask', response.data)
                });
        }
    },
    mutations: {
        addTask (state, task) {
            state.tasks.push(task)
        },
        setTasks (state, tasks) {
            state.tasks = tasks
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
