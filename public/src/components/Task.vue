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
                input: '',
                output: '',
                isUploaded: false
            }
        },
        computed: {
            ...mapGetters([
                    'tasks',
                    'taskCount'
            ]),
            task(){
                return this.$store.getters.getTaskById(this.id)
            }
        },
        methods: {
            upload() {
                let data = new FormData();
                data.append('file', this.file);
                axios.post('/upload', data, {
                    headers: {
                        'Content-Type': 'multipart/form-data'
                    }
                });
            },
            launch() {
                console.log("Launched")
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