<template>
    <div id="app">
        <div class="row">
            <div class="col-md-3">
                <navigation></navigation>
            </div>
            <div class="col-md-9">
                <router-view></router-view>
            </div>
        </div>
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
            let files = this.$refs.uploadfile.files;
            let data = new FormData();
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
    },
    created() {
        this.$store.dispatch('fetchTasks');
        this.$store.dispatch('fetchImages')
    }
}
</script>
