<template>
        <div>
                <v-data-table :headers="scoreBoard.headers" :loading="table.loading" :items="scoreBoard.data"
                ></v-data-table>
        </div>
</template>


<script>
    import * as _ from 'lodash';
    export default {
        data() {
            return {
                dialog: {
                    show: false,
                    data: {
                        name: '',
                    }
                },
                table: {
                    loading: false,
                    headers: [
                        {text: 'Takım İsmi', value: 'name'},
                        {text: 'Oynadığı Lig', value: 'league.name'},
                        {text: 'Güncelleme', value: 'updated_at'},
                        {text: 'Aksiyonlar', value: 'actions', sortable: false},
                    ],
                    data: []
                },
                scoreBoard: {
                    data: [],
                    headers: [
                        {text: 'Takım İsmi', value: 'team.name'},
                        {text: 'G', value: 'won'},
                        {text: 'B', value: 'drawn'},
                        {text: 'M', value: 'lost'},
                        {text: 'A', value: 'for'},
                        {text: 'Y', value: 'against'},
                        {text: 'AV', value: 'goal_diff'},
                        {text: 'P', value: 'points'},
                    ]
                },
            }
        },
        computed: {
            teamId() {
                return this.$route.params.team_id
            },
            _() {
                return _;
            }
        },
        mounted() {
            this.init()
        },
        methods: {
            /**
             * @returns {void}
             */
            async init() {
                await this.refresh();
                this.table.loading = true;
                const {data} = await this.$api.team.get(this.teamId);
                this.scoreBoard.data = data;
                const predict = await this.$api.match.predict(this.leagueId);
                this.predict = predict.data;
                this.table.loading = false;
            },

            /**
             * @returns {void}
             */
            async save() {
                this.dialog.show = false;
                this.table.loading = true;
                if (this.dialog.data.id) {
                    await this.$api.team.update(this.leagueId, this.dialog.data.id, {name: this.dialog.data.name});
                } else {
                    await this.$api.team.create(this.leagueId, {name: this.dialog.data.name})
                }
                await this.close();
                await this.refresh();
            },

            /**
             * @param {object} item
             * @returns {void}
             */
            async editItem(item) {
                this.dialog.data = item;
                this.dialog.show = true;
            },

            /**
             * @param {number} id
             * @returns {void}
             */
            async deleteItem(id) {
                this.table.loading = true;
                await this.$api.team.delete(this.leagueId, id);
                await this.refresh();
            },

            /**
             * @returns {void}
             */
            async refresh() {
                this.table.loading = true;
                const {data} = await this.$api.team.list(this.leagueId);
                this.table.data = data;
                this.table.loading = false;
            }
        }
    }
</script>
