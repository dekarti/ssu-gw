<template>
    <div class="form-main">
        <div class="form-field">
            <form-input
                    type="text"
                    label="Task name"
                    :value="name"
                    placeholder="Finite machine">
            </form-input>
        </div>
        <div class="form-container">
            <div class="form-area">
                <label class="form-label">Description</label>
                <div>
                    <textarea cols="45"
                              rows="20"
                              class="form-input"
                              placeholder="# Task 1"
                              @input="update">{{ description }}</textarea>
                </div>
            </div>
            <div class="form-area">
                <label class="form-label">Formatted Description</label>
                <div v-html="compiledMarkdown"></div>
            </div>
        </div>

        <hr/>

        <div class="form-container">
            <div class="form-container__part">
                <div>
                    <label class="form-label">Test Input</label>
                </div>
                <div class="form-area">
                    <div>
                        <textarea cols="45"
                                  rows="20"
                                  class="form-input">{{ defaultInput }}</textarea>
                    </div>
                </div>
            </div>
            <div class="form-container__part">
                <div>
                    <label class="form-label">Expected Output</label>
                </div>
                <div class="form-area">
                    <div>
                        <textarea cols="45"
                                  rows="20"
                                  class="form-input">{{ expectedOutput }}</textarea>
                    </div>
                </div>
            </div>
        </div>
        <div class="button">
                <button @click="add">Launch</button>
        <div>
    </div>
<!--        <div class="form-item">
            <label class="form-label">Task name</label>
            <input class="form-input" v-model="name" type="text" placeholder="Finite machine" />
        </div>
        <div class="editor">
            <label class="form-label">Raw Description</label>
            <textarea :value="description" @input="update"></textarea>
        </div>
        <div class="editor">
            <label class="form-label">Input</label>
            <textarea v-model="defaultInput"></textarea>
        </div>
        <div class="editor">
            <label class="form-label">Output</label>
            <textarea v-model="expectedOutput"></textarea>
        </div>-->
    </div>
</template>

<script>
    import _ from 'underscore';
    import marked from 'marked';
    import FormInput from './FormInput.vue'
    import FormTextArea from './FormTextArea.vue'

    export default {
        components: { FormInput, FormTextArea },
        data() {
            return {
                name: '',
                description: '',
                defaultInput: '',
                expectedOutput: ''
            }
        },
        computed: {
            tasks() {
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
            },
            compiledMarkdown: function () {
                return marked(this.description, {sanitize: true})
            }
        },
        methods: {
            update: _.debounce(function (e) {
                    this.description = e.target.value
            }, 300),
            add() {

            }
        }
    }
</script>

<style>
    @import "~variables/variables.scss";
    @import "~variables/media-queries.scss";

    .form-field {
        width: 70%;
        margin-top: 10px;
        margin-bottom: 40px;
    }
    .form-main {
        width: 100%;
        margin-left: 10px;
        margin-right: 10px;
    }

    .form-container {
        width: 100%;
        height: 300px;
    }
    .form-container__part {
        float: left;
        width: 40%;
        height: 90%;
    }
    .form-area {
        width: 50%;
        height: 100%;
        float: left;
    }
    .button {
        height: 36px;
        line-height: 32px;
        background: none;
        padding: 0 1rem;
        overflow: visible;
        border-radius: 5px;
        margin-left: 1rem;
        letter-spacing: 1px;
        @include font-size(13);
        color: $color-disabled;
        text-transform: uppercase;
        border: 2px solid $color-border;
        font-weight: $font-weight-semibold;
        @include transition (all 0.12s linear);
        float: right;
    }
</style>

