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

<style scoped>
    .form-item {
        position: relative;
        margin-bottom: 0.7rem;
    }

    .form-label {
    @include text-label();
        display: inline-block;
        margin-bottom: 1px;
    }

    .form-help {
        color: $ color-ok-text;
    }

    .form-input {
        width: 100%;
        height: 40px;
        padding: 0 0.7rem;
    @include text-main();
        border: 2px solid $ color-border;
    @include transition (all 0.12 s linear);

    &
    :hover {
        border: 2px solid darken($ color-border, 15);
    }

    &
    :focus {
        border: 2px solid $ color-ok-solid;
    }

    &
    ::-webkit-input-placeholder { /* WebKit, Blink, Edge */
    @include text-placeholder();
    }

    &
    :-moz-placeholder { /* Mozilla Firefox 4 to 18 */
    @include text-placeholder();
        opacity: 1;
    }

    &
    ::-moz-placeholder { /* Mozilla Firefox 19+ */
    @include text-placeholder();
        opacity: 1;
    }

    &
    :-ms-input-placeholder { /* Internet Explorer 10-11 */
    @include text-placeholder();
    }

    &
    :placeholder-shown { /* Standard (https://drafts.csswg.org/selectors-4/#placeholder) */
    @include text-placeholder();
    }

    }

    .form-select {
        width: 100%;
        height: 40px;
        padding: 0 0.7rem;
        line-height: 36px;
    @include text-main();
        border: 2px solid $ color-border;
        position: relative;
        z-index: 30;
    @include transition (all 0.12 s linear);

    &
    :placeholder-shown {
    @include text-placeholder();
    }

    &
    :hover {
        border: 2px solid darken($ color-border, 15);
        cursor: pointer;

    .form-select-dropdown {
        border: 2px solid darken($ color-border, 15);
    }

    .dropdown-caret {
        color: darken($ color-border, 35);
    }

    .dropdown-caret {
        color: $ color-ok-text;
    }

    }
    .dropdown-caret {
        position: absolute;
        right: 10px;
        top: 1px;
        color: darken($ color-border, 20);
    }

    }

    .form-select-dropdown {
        position: absolute;
        top: 63px;
        background: $ color-white;
        border: 2px solid $ color-border;
        border-top: 0;
        width: 88%;
        overflow: scroll;
        z-index: 20;
        max-height: 0px;
    @include transition (all 0.12 s ease-in);

    li {
        padding: 0 10px;
        line-height: 36px;
    @include transition (all 0.06 s linear);

    &
    :hover {
        cursor: pointer;
        background: lighten($ color-border, 17);
        color: $ color-title;
    }

    }
    }

    .open {

    .form-select {
        border: 2px solid $ color-ok-solid;
        border-bottom: 0;
    }

    .form-select-dropdown {
        border: 2px solid $ color-ok-solid;
        border-top: 0;
        max-height: 200px;
    }

    .dropdown-caret {
        color: $ color-ok-text;
    }

    }

    .form-error {

    .form-label {
    @include text-label();
        color: $ color-warning-text;
    }

    .form-input {
        border: 2px solid $ color-warning-solid;
    }

    }

    label.form-checkbox:hover {
        cursor: pointer;

    span {
        border: 2px solid darken($ color-border, 15);
    }

    }

    .form-checkbox {
        height: 20px;
        line-height: 20px;
        display: inline-block;
    @include transition (all 0.12 s linear);

    span {
        width: 20px;
        height: 20px;
        display: block;
        background: $ color-white;
        border: 2px solid $ color-border;
    @include transition (all 0.12 s linear);
        float: left;
        margin-right: 0.5rem;
    }

    input {
        display: none;
    }

    }
    .form-checkbox-ghost .icon-tick {
        height: 16px;
        width: 17px;
        line-height: 17px;
        text-align: center;
        display: block;
        font-size: 13px;
    }
</style>