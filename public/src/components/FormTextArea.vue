<template>
    <div class="form-item" :class="{'form-error' : formError}">
        <label class="form-label" v-if="!formError">{{ label }}</label>
        <label class="form-label" v-if="formError">{{ message }}</label>
        <div>
            <textarea cols="100" rows="15">{{ value }}</textarea>
        </div>
    </div>
</template>

<script>
    export default {

        props: ['type', 'label', 'error', 'value', 'message', 'placeholder'],

        data() {
            return {
                errorMessage: 'This is an error message'
            }
        },

        computed: {
            inputText() {
                return {
                    'type': this.type,
                    'label': this.label,
                    'error': this.error,
                    'value': this.value,
                    'message': this.message,
                    'placeholder': this.placeholder
                }
            },
            formError() {
                if (this.error === 'true') {
                    return true
                }
            }
        },
    };
</script>

