<template>
    <div class="navigation">
        <b-nav vertical pills>
    <!--        <router-link :to="{name: 'home'}">
                <img src="~assets/logo.png">
            </router-link>-->
            <b-nav-item>
                <router-link class="nav__link" :tasks=tasks :to="{name: 'new'}">New Task</router-link>
            </b-nav-item>
            <b-nav-item v-for="task in tasks">
                <router-link tasks="tasks" :to="{name: 'task', params: {id: task.id}}">{{ task.name }}</router-link>
            </b-nav-item>
        </b-nav>
    </div>
</template>

<script>
    import { mapGetters } from 'vuex'

    export default {
        computed: {
            ...mapGetters([
                'tasks',
                'taskCount'
            ]),
        },
        methods: {
            toggleNav(){
                document.querySelector('.nav__hamburger').classList.toggle('nav__hamburger--active');
                document.querySelector('.nav__list').classList.toggle('nav__list--active');
            }
        },
    }
</script>

<style>
</style>
