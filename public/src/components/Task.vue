<template>
    <div>
        <h1>{{ task.name }}</h1>
        <p>{{ task.description }}</p>
    </div>
</template>

<script>
    export default {
        props: ['id'],
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
            },
            task(){
                return this.tasks.find(x => x.id === this.id)
            }
        }
    }
</script>

<style>

</style>