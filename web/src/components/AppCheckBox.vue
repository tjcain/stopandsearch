<template>
    <div>
        <p class="menu-label">{{listBoxes.title}}</p>
        <section class="checkboxes" @input="emitValues()">
            <div v-for="box in listBoxes.parameters" :key=box.key class="field">
                <b-checkbox 
                    size="is-small"
                    v-model="boxes.values"
                    :native-value="box.value">
                        {{box.label}}
                </b-checkbox>
            </div>
        </section>
        {{boxes.values}}
    </div>
</template>

<script>
export default {
    props: {
        boxes: {
            required: true
        }
    },
    computed: {
        listBoxes() {
            return this.boxes
        },
    },
    methods: {
        emitValues: function() {
            this.$emit('change-values', this.boxes.values)
        },
    },
}
</script>

<style scoped>
    .field:last-child {
    margin-bottom: 1px;
    }
    .field {
    margin: 1px 0px;
    }
    .checkboxes {
    max-height: 15vh;
    overflow-y: scroll;
    }
    .menu-label {
    margin: 1px 0
    }

</style>
