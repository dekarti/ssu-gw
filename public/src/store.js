import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex);

const store = new Vuex.Store({
    state: {
        tasks: [
            {
                id: 1,
                name: "task1",
                description: "task1"
            }
        ],
        images: [
            {
                name: "python",
                tags: [
                    "2.7",
                    "2.7-slim"
                ]
            },
            {
                name: "java",
                tags: [
                    "7",
                    "8"
                ]
            }
        ]
    },
    actions: {
        fetchTasks ({ commit }) {
            axios.get("/tasks")
                .then((response) => {
                    commit('setTasks', response.data.tasks)
                })
        },
        fetchImages ({ commit }) {
            axios.get("/images")
                .then((response) => {
                    commit('setImages', response.data.images)
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
        },
        setImages (state, images) {
            state.images = images
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
            },
            images: state => {
                return state.images
            },
            filterImagesByName: (state) => (name) => {
                return state.images.find(x => x.name.startWith(name))
            }
    }
});

export default store
