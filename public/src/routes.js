import VueRouter from 'vue-router';

let routes = [
    {
        name: 'home',
        path: '/',
        components: {
        }
    },
    {
        name: 'new',
        path: '/task/new',
        component: require('./components/NewTask.vue')
    },
    {
        name: 'task',
        path: '/task/:id',
        component: require('./components/TaskPage.vue')
    }
];

const router = new VueRouter({
    routes
});

export default router;
