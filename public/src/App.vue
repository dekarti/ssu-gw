<template>
    <div id="app">
        <navigation></navigation>
        <header class="header"></header>
<!--        <textarea cols="40" rows="15" charswidth="23" v-model="input"></textarea>
        <textarea cols="40" rows="15" charswidth="23" v-model="output"></textarea>
        <input ref="uploadfile" type="file" @change="upload">
        <button @click="launch">Launch</button>-->
        <section class="main">
            <router-view></router-view>
        </section>
    </div>
</template>

<script>
import axios from 'axios'
import Navigation from './components/Navigation.vue'

export default {
    name: 'app',
    components: { Navigation },
    data () {
        return {
            input: '1e-3\n116',
            output: '',
            expectedOutput: 'True\nTrue',
            currentTask: 0
        }
    },
    methods: {
        upload: function (e) {
            e.preventDefault();
            var files = this.$refs.uploadfile.files;
            var data = new FormData();
            data.append('file', files[0]);
            axios.post('/upload', data, {
                headers: {
                    'Content-Type': 'multipart/form-data'
                }
            });
        },
        launch: function (e) {
            e.preventDefault();
            self = this;
            axios.post('/task/' + this.currentTask + '/launch', {
                input: this.input
            })
                .then(function (response) {
                    console.log(response);
                    var escapedResponse = JSON
                        .stringify(response)
                        .replace(/\\n/g, "\\n")
                        .replace(/\\'/g, "\\'")
                        .replace(/\\"/g, '\\"')
                        .replace(/\\&/g, "\\&")
                        .replace(/\\r/g, "\\r")
                        .replace(/\\t/g, "\\t")
                        .replace(/\\b/g, "\\b")
                        .replace(/\\f/g, "\\f");
                    self.output = JSON.parse(escapedResponse).data.output;
                    console.log(response);
                })
                .catch(function (error) {
                    console.log(error);
                });
        }
    }
}
</script>

<style lang="scss">
    @import "./src/scss/variables";
    @import "./src/scss/media-queries";

    *{
        box-sizing: border-box;
    }
    html, body{
        height: 100%;
    }
    body{
        font-family: 'Roboto', sans-serif;
        line-height: 1.6;
        background: $c-light;
        color: $c-dark;
    &.hidden{
         overflow: hidden;
     }
    }
    input, textarea, button{
        font-family: 'Roboto', sans-serif;
    }
    figure{
        padding: 0;
        margin: 0;
    }
    img{
        display: block;
        max-width: 100%;
        height: auto;
    }
    .loader{
        animation: load 1s linear infinite;
        border: 2px solid $c-white;
        border-radius: 50%;
        display: block;
        height: 30px;
        left: 50%;
        margin: -1.5em;
        position: absolute;
        top: 50%;
        width: 30px;
    &:after {
         border: 5px solid $c-green;
         border-radius: 50%;
         content: '';
         left: 10px;
         position: absolute;
         top: 16px;
     }
    }
    @keyframes load {
        100% { transform: rotate(360deg); }
    }
    .wrapper{
        position: relative;
    }
    .header{
        position: fixed;
        background: $c-white;
        z-index: 15;
        display: flex;
    @include tablet-min{
        width: calc(100% - 170px);
        height: 75px;
        margin-left: $nav-width;
        border-top: 0;
        border-bottom: 0;
        top: 0;
    }
    &__search{
         height: 50px;
         display: flex;
         position: relative;
         z-index: 5;
         width: calc(100% - 110px);
         position: fixed;
         top: 0;
         right: 55px;
    @include tablet-min{
        position: relative;
        height: 75px;
        right: 0;
    }
    &-input{
         display: block;
         width: 100%;
         padding: 15px 20px 15px 45px;
         outline: none;
         border: 0;
         background-color: transparent;
         color: $c-dark;
         font-weight: 300;
         font-size: 16px;
    @include tablet-min{
        padding: 15px 30px 15px 60px;
    }
    @include tablet-landscape-min{
        padding: 15px 30px 15px 80px;
    }
    @include desktop-min{
        padding: 15px 30px 15px 90px;
    }
    }
    &-icon{
         width: 14px;
         height: 14px;
         fill: rgba($c-dark, 0.5);
         transition: fill 0.5s ease;
         pointer-events: none;
         position: absolute;
         top: 50%;
         margin-top: -7px;
         left: 20px;
    @include tablet-min{
        width: 18px;
        height: 18px;
        margin-top: -9px;
        left: 30px;
    }
    @include tablet-landscape-min{
        left: 50px;
    }
    @include desktop-min{
        left: 60px;
    }
    }
    &-input:focus + &-icon{
         fill: $c-dark;
     }
    }
    }
    .main{
        position: relative;
        padding: 50px 0 0;
    @include tablet-min{
        width: calc(100% - #{$nav-width});
        padding: 75px 0 0;
        margin-left: $nav-width;
        position: relative;
    }
    }
    .button{
        display: inline-block;
        border: 1px solid $c-dark;
        text-transform: uppercase;
        background: $c-dark;
        font-weight: 300;
        font-size: 11px;
        line-height: 2;
        letter-spacing: 0.5px;
        padding: 5px 20px 4px 20px;
        cursor: pointer;
        color: $c-dark;
        background: transparent;
        outline: none;
        transition: background 0.5s ease, color 0.5s ease;
    @include tablet-min{
        font-size: 12px;
        padding: 6px 20px 5px 20px;
    }
    body:not(.touch) &:hover{
                          background: $c-dark;
                          color: $c-white;
                      }
    }
    // router view transition
       .fade-enter-active, .fade-leave-active {
           transition-property: opacity;
           transition-duration: 0.25s;
       }
    .fade-enter-active {
        transition-delay: 0.25s;
    }
    .fade-enter, .fade-leave-active {
        opacity: 0
    }

    textarea {
        resize: none;
    }
</style>
