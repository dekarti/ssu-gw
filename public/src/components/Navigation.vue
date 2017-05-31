<template>
    <nav class="nav">
        <router-link class="nav__logo" :to="{name: 'home'}">
            <img class="nav__logo-image" src="~assets/logo.png">
        </router-link>
        <div class="nav__hamburger" @click="toggleNav">
            <div class="bar"></div>
            <div class="bar"></div>
            <div class="bar"></div>
        </div>
        <ul class="nav__list">
            <li>
                <router-link class="nav__link" :to="{name: 'new'}">
                    <div class="nav__link-wrap">
                        <span class="nav__link-title">New Task</span>
                    </div>
                </router-link>
            </li>
            <li v-for="task in tasks">
                <router-link class="nav__link" :to="{name: 'task', params: {id: task.id}}">
                    <div class="nav__link-wrap">
                        <span class="nav__link-title">{{ task.name }}</span>
                    </div>
                </router-link>
            </li>
        </ul>
    </nav>
</template>

<script>
    import NewTask from './NewTask.vue'

    export default {
        components: { NewTask },
        computed: {
            tasks(){
                return [
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
            }
        },
        methods: {
            toggleNav(){
                document.querySelector('.nav__hamburger').classList.toggle('nav__hamburger--active');
                document.querySelector('.nav__list').classList.toggle('nav__list--active');
            }
        },
    }
</script>

<style lang="scss">
    @import "./src/scss/variables";
    @import "./src/scss/media-queries";

    .nav{
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 50px;
        background: $c-white;
        display: flex;
        z-index: 10;
        @include tablet-min{
            display: block;
            width: $nav-width;
            height: 100vh;
        }
        &__logo{
            width: $nav-width;
            height: 75px;
            display: flex;
            align-items: center;
            justify-content: center;
            background: $c-dark;
            &-image{
                width: 55px;
                height: 50px;
                fill: $c-green;
                transition: transform 0.5s ease;
            }
            &:hover &-image{
                transform: scale(1.04);
            }
        }
        &__hamburger{
            display: none;
            position: fixed;
            width: 55px;
            height: 50px;
            top: 0;
            right: 0;
            cursor: pointer;
            background: $c-white;
            z-index: 10;
            border-left: 1px solid $c-light;
            .bar{
                position: absolute;
                width: 23px;
                height: 1px;
                background: rgba($c-dark, 0.5);
                transition: all 300ms ease;
                &:nth-child(1){
                    left: 16px;
                    top: 17px;
                }
                &:nth-child(2){
                    left: 16px;
                    top: 25px;
                    &:after {
                        content: "";
                        position: absolute;
                        left: 0px;
                        top: 0px;
                        width: 23px;
                        height: 1px;
                        background: transparent;
                        transition: all 300ms ease;
                    }
                }
                &:nth-child(3){
                    right: 15px;
                    top: 33px;
                }
            }
            &--active{
                .bar{
                    &:nth-child(1),
                    &:nth-child(3){
                        width: 0;
                    }
                    &:nth-child(2) {
                        transform: rotate(-45deg);
                    }
                    &:nth-child(2):after {
                        transform: rotate(-90deg);
                        background: rgba($c-dark, 0.5);
                    }
                }
            }
        }
        &__list{
            list-style: none;
            padding: 0;
            margin: 0;
            text-align: center;
            left: 0;
            border-top: 1px solid $c-light;
            background: transparent;
            position: relative;
            display: block;
            width: 100%;
            border-top: 0;
            top: 0;
        }
        &__link{
            width: 100%;
            height: 45px;
            display: flex;
            flex-wrap: wrap;
            align-items: center;
            justify-content: left;
            margin-left: 20px;
            font-size: 12px;
            font-weight: 300;
            text-decoration: none;
            text-transform: uppercase;
            letter-spacing: 0.5px;
            color: rgba($c-dark, 0.7);
            transition: color 0.5s ease, background 0.5s ease;
            position: relative;
            cursor: pointer;
            &-title{
                display: block;
                width: 100%;
            }
            &:hover{
                color: $c-dark;
            }
            &:hover &-icon{
                fill: $c-dark;
            }
            &.is-active{
                color: $c-dark;
                background: $c-light;
            }
            &.is-active &-icon{
                fill: $c-dark;
            }
        }
    }
</style>