<template>
    <div>
        <div v-html="task.description"></div>

        <br/>

        <div class="row">
            <div class="col-md-6">
                <b-form-file v-model="file" @input="upload"></b-form-file>
                <br> Selected file: {{file && file.name}}
            </div>
            <div class="col-md-6">
                <b-button variant="primary" @click="launch">Launch</b-button>
            </div>
        </div>

        <div class="row">
            <div class="col-sm-6">
                <small class="text-muted">Command to execute</small>
                <b-form-input v-model="command"
                              type="text"
                              placeholder="python main.py"></b-form-input>
            </div>
            <div class="col-sm-3">
                <small class="text-muted">Language</small>
                <select class="col-sm-12" style="height: 37px" v-model="language">
                    <option disabled value="">Please select one</option>
                    <option v-for="image in images">{{ image.name }}</option>
                </select>
            </div>
            <div class="col-sm-3">
                <small class="text-muted">Version</small>
                <select class="col-sm-12" style="height: 37px" v-model="version">
                    <option disabled value="">Please select one</option>
                    <option v-for="tag in currentImage.tags">{{ tag }}</option>
                </select>
            </div>
        </div>

        <div class="row">
            <div class="col-md-6">
                <small class="text-muted">Input</small>
                <b-form-input class="area"
                              textarea
                              v-model="input"
                              type="text"></b-form-input>
            </div>
            <div class="col-md-6">
                <small class="text-muted">Output</small>
                <b-form-input class="area"
                              textarea
                              v-model="output"
                              type="text"></b-form-input>
            </div>
        </div>
    </div>
</template>

<script>
    import { mapGetters } from 'vuex'
    import axios from 'axios'

    export default {
        props: ['id'],
        data() {
            return {
                file: null,
                command: '',
                language: '',
                version: '',
                input: '',
                output: '',
                path: '',
                isUploaded: false,
                currentImage: {}
            }
        },
        computed: {
            ...mapGetters([
                    'tasks',
                    'taskCount',
                    'images'
            ]),
            task() {
                return this.$store.getters.getTaskById(this.id)
            },
        },
        methods: {
            upload() {
                let data = new FormData();
                data.append('file', this.file);
                axios.post('/upload', data, {
                    headers: {
                        'Content-Type': 'multipart/form-data'
                    }
                }).then((resp) => {
                    this.path = resp.data.path;
                });
            },
            launch() {
                let data = {
                    task: this.task,
                    config: {
                        command: this.command.split(" "),
                        language: this.language,
                        version: this.version
                    },
                    input: this.input,
                    path: this.path
                };
                axios.post('/tasks/launch', data)
                    .then((resp) => {
                        this.output = resp.data.output;
                        this.input = resp.data.input;
                    });
            }
        },
        watch: {
            language: function() {
                for (var i in this.images) {
                    if (this.images[i].name === this.language) {
                        this.currentImage = this.images[i]
                    }
                }
            }
        }

    }
</script>


<style>
    .custom-file-control::after {
        content: 'No files selected' !important;
    }

    .custom-file-control::before {
        content: 'Choose file' !important;
    }

    .custom-file .drop-here::before {
        content: 'Drop files here' !important;
    }
</style>