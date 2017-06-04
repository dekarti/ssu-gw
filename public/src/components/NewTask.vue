<template>
    <div class="main">
        <div class="header"></div>

        <div>
            <small class="text-muted">Enter task name</small>
            <div class="row">
                <div class="col-sm-8">
                    <b-form-input v-model="name"
                                  type="text"
                                  placeholder="Finite machine"></b-form-input>
                </div>
                <div class="col-sm-4">
                    <b-button variant="primary" @click="add">Add</b-button>
                </div>
            </div>

            <br/>
        </div>

        <div class="row">
            <div class="col-md-6">
                <small class="text-muted">Enter task description (Markdown is available for formatting)</small>
                <b-form-input class="area"
                              textarea
                              v-model="description"
                              type="text"
                              placeholder="# Task 1"></b-form-input>
            </div>

            <div class="col-md-6">
                <small class="text-muted">Formatted description</small>
                <div class="area" v-html="compiledMarkdown"></div>
            </div>
        </div>

        <br/>

        <div class="row">
            <div class="col-md-6">
                <small class="text-muted">Input</small>
                <b-form-input class="area"
                              textarea
                              v-model="defaultInput"
                              type="text"></b-form-input>
            </div>
            <div class="col-md-6">
                <small class="text-muted">Expected Output</small>
                <b-form-input class="area"
                              textarea
                              v-model="expectedOutput"
                              type="text"></b-form-input>
            </div>
        </div>

<!--        <div>
            <button @click="add">Launch</button>
        </div>-->
    </div>
</template>

<script>
    import _ from 'underscore';
    import marked from 'marked';
    import FormInput from './FormInput.vue'
    import FormTextArea from './FormTextArea.vue'
    import { mapGetters, mapActions } from 'vuex'


    export default {
        components: { FormInput, FormTextArea },
        data() {
            return {
                name: '',
                description: '',
                defaultInput: '',
                expectedOutput: '',
                submitted: false
            }
        },
        computed: {
            ...mapGetters([
                'tasks',
                'taskCount'
            ]),
            compiledMarkdown: function () {
                return marked(this.description, {sanitize: true})
            }
        },
        methods: {
            update: _.debounce(function (e) {
                    this.description = e.target.value
            }, 300),

            add() {
                this.$store.dispatch('addTask', {
                    name: this.name,
                    description: this.compiledMarkdown,
                    defaultInput: this.defaultInput,
                    expectedOutput: this.expectedOutput
                })
            }
        }
    }
</script>

<style>
/*    .main {
        float: left;
    }

    .form {
       width: 700px;
    }
    .header {
        width: 100%;
        height: 15px;
    }

    .block {
        width: 400px;
        height: 300px;
        float: left;
        margin-right: 50px;
    }
    .area {
        width: 300px;
        height: 250px;
    }*/
.area {
    height: 250px;
}
</style>
