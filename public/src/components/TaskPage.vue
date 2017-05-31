<template>
    <task :id="$route.params.id"></task>
</template>

<script>
    import Task from './Task.vue';
    export default {
        components: { Task }
    }
</script>