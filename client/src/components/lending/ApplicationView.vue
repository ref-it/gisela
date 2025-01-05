<template>
    <div class="p-6 sm:p-8 overflow-y-auto">
        <div class="divide-y divide-zinc-200 dark:divide-zinc-600 overflow-hidden rounded-lg bg-zinc-100 dark:bg-zinc-700 text-zinc-800 dark:text-white shadow ring-1 ring-zinc-200 dark:ring-zinc-600 mb-8">
            <div class="px-4 py-3 uppercase font-bold">
                Status
            </div>
            <div class="p-4 overflow-x-auto">
                <svg ref="applicationFlow" class="w-[384px] h-[534px] lg:w-[1044px] lg:h-[154px] mx-auto">
                    <g class="translate-x-[2px] translate-y-[2px]">
                        <rect width="160" height="50" rx="6" ry="6" class="fill-zinc-200 dark:fill-zinc-800 stroke-[1.5] stroke-zinc-600 dark:stroke-zinc-400" />
                        <text x="80" y="29" class="fill-zinc-800 dark:fill-white" text-anchor="middle">Eingereicht</text>

                        <path d="M 161,25 T 214,25" class="stroke-[1.5] stroke-zinc-600 dark:stroke-zinc-400 fill-none"></path>
                        <path d="M 216,25 L 206,22 V 28 Z" class="stroke-[1.5] stroke-zinc-600 dark:stroke-zinc-400 fill-zinc-600 dark:fill-zinc-400"></path>

                        <path d="M 80,51 T 80,95" class="stroke-[1.5] stroke-zinc-600 dark:stroke-zinc-400 fill-none"></path>
                        <path d="M 80,96 L 77,86 H 83 Z" class="stroke-[1.5] stroke-zinc-600 dark:stroke-zinc-400 fill-zinc-600 dark:fill-zinc-400"></path>
                    </g>
                    <g class="translate-x-[222px] translate-y-[2px] lg:translate-x-[2px] lg:translate-y-[102px]">
                        <rect width="160" height="50" rx="6" ry="6" class="fill-zinc-200 dark:fill-zinc-800 stroke-[1.5] stroke-zinc-600 dark:stroke-zinc-400" />
                        <text x="80" y="29" class="fill-zinc-800 dark:fill-white" text-anchor="middle">Zurückgezogen</text>
                    </g>
                    <g class="translate-x-[2px] translate-y-[102px] lg:translate-x-[222px] lg:translate-y-[2px]">
                        <rect width="160" height="130" rx="6" ry="6" class="fill-zinc-200 dark:fill-zinc-800 stroke-[1.5] stroke-zinc-600 dark:stroke-zinc-400" />
                        <text x="80" y="29" class="fill-zinc-800 dark:fill-white" text-anchor="middle">In Bearbeitung</text>

                        <path d="M 161,25 T 214,25" class="stroke-[1.5] stroke-zinc-600 dark:stroke-zinc-400 fill-none"></path>
                        <path d="M 216,25 L 206,22 V 28 Z" class="stroke-[1.5] stroke-zinc-600 dark:stroke-zinc-400 fill-zinc-600 dark:fill-zinc-400"></path>

                        <path d="M 161,105 T 214,105" class="hidden lg:inline stroke-[1.5] stroke-zinc-600 dark:stroke-zinc-400 fill-none"></path>
                        <path d="M 216,105 L 206,102 V 108 Z" class="hidden lg:inline stroke-[1.5] stroke-zinc-600 dark:stroke-zinc-400 fill-zinc-600 dark:fill-zinc-400"></path>

                        <path d="M 80,131 T 80,175" class="lg:hidden stroke-[1.5] stroke-zinc-600 dark:stroke-zinc-400 fill-none"></path>
                        <path d="M 80,176 L 77,166 H 83 Z" class="lg:hidden stroke-[1.5] stroke-zinc-600 dark:stroke-zinc-400 fill-zinc-600 dark:fill-zinc-400"></path>

                        <g transform="translate(20,45)">
                            <rect id="ab-ok" width="120" height="30" rx="6" ry="6" class="fill-zinc-200 dark:fill-zinc-800 stroke-[1.5] stroke-zinc-600 dark:stroke-zinc-400" />
                            <text x="60" y="21" class="fill-zinc-800 dark:fill-white" text-anchor="middle">AB-OK</text>
                        </g>

                        <g transform="translate(20,85)">
                            <rect id="hv-ok" width="120" height="30" rx="6" ry="6" class="fill-zinc-200 dark:fill-zinc-800 stroke-[1.5] stroke-zinc-600 dark:stroke-zinc-400" />
                            <text x="60" y="21" class="fill-zinc-800 dark:fill-white" text-anchor="middle">HV-OK</text>
                        </g>
                    </g>
                    <g class="translate-x-[222px] translate-y-[102px] lg:translate-x-[442px] lg:translate-y-[82px]">
                        <rect id="rejected" width="160" height="50" rx="6" ry="6" class="fill-zinc-200 dark:fill-zinc-800 stroke-[1.5] stroke-zinc-600 dark:stroke-zinc-400" />
                        <text x="80" y="29" class="fill-zinc-800 dark:fill-white" text-anchor="middle">Abgelehnt</text>
                    </g>
                    <g class="translate-x-[2px] translate-y-[282px] lg:translate-x-[442px] lg:translate-y-[2px]">
                        <rect id="ready" width="160" height="50" rx="6" ry="6" class="fill-zinc-200 dark:fill-zinc-800 stroke-[1.5] stroke-zinc-600 dark:stroke-zinc-400" />
                        <text x="80" y="29" class="fill-zinc-800 dark:fill-white" text-anchor="middle">Bereit</text>

                        <path d="M 161,25 T 214,25" class="hidden lg:inline stroke-[1.5] stroke-zinc-600 dark:stroke-zinc-400 fill-none"></path>
                        <path d="M 216,25 L 206,22 V 28 Z" class="hidden lg:inline stroke-[1.5] stroke-zinc-600 dark:stroke-zinc-400 fill-zinc-600 dark:fill-zinc-400"></path>

                        <path d="M 80,51 T 80,95" class="lg:hidden stroke-[1.5] stroke-zinc-600 dark:stroke-zinc-400 fill-none"></path>
                        <path d="M 80,96 L 77,86 H 83 Z" class="lg:hidden stroke-[1.5] stroke-zinc-600 dark:stroke-zinc-400 fill-zinc-600 dark:fill-zinc-400"></path>
                    </g>
                    <g class="translate-x-[2px] translate-y-[382px] lg:translate-x-[662px] lg:translate-y-[2px]">
                        <rect id="borrowed" width="160" height="50" rx="6" ry="6" class="fill-zinc-200 dark:fill-zinc-800 stroke-[1.5] stroke-zinc-600 dark:stroke-zinc-400" />
                        <text x="80" y="29" class="fill-zinc-800 dark:fill-white" text-anchor="middle">Ausgeliehen</text>

                        <path d="M 161,25 T 214,25" class="hidden lg:inline stroke-[1.5] stroke-zinc-600 dark:stroke-zinc-400 fill-none"></path>
                        <path d="M 216,25 L 206,22 V 28 Z" class="hidden lg:inline stroke-[1.5] stroke-zinc-600 dark:stroke-zinc-400 fill-zinc-600 dark:fill-zinc-400"></path>

                        <path d="M 80,51 T 80,95" class="lg:hidden stroke-[1.5] stroke-zinc-600 dark:stroke-zinc-400 fill-none"></path>
                        <path d="M 80,96 L 77,86 H 83 Z" class="lg:hidden stroke-[1.5] stroke-zinc-600 dark:stroke-zinc-400 fill-zinc-600 dark:fill-zinc-400"></path>
                    </g>
                    <g class="translate-x-[2px] translate-y-[482px] lg:translate-x-[882px] lg:translate-y-[2px]">
                        <rect id="returned" width="160" height="50" rx="6" ry="6" class="fill-zinc-200 dark:fill-zinc-800 stroke-[1.5] stroke-zinc-600 dark:stroke-zinc-400" />
                        <text x="80" y="29" class="fill-zinc-800 dark:fill-white" text-anchor="middle">Zurückgegeben</text>
                    </g>

                    <g class="translate-x-[222px] translate-y-[442px] lg:translate-x-[712px] lg:translate-y-[112px]">
                        <rect id="" width="160" height="40" rx="6" ry="6" class="fill-green-600 stroke-[1.5] stroke-green-700" />
                        <text x="80" y="24" class="fill-white" text-anchor="middle">Aktueller Status</text>
                    </g>
                    <g class="translate-x-[222px] translate-y-[492px] lg:translate-x-[882px] lg:translate-y-[112px]">
                        <rect id="" width="160" height="40" rx="6" ry="6" class="fill-sky-600 stroke-[1.5] stroke-sky-700" />
                        <text x="80" y="24" class="fill-white" text-anchor="middle">Wechsel möglich</text>
                    </g>
                </svg>
            </div>
        </div>

        <div class="divide-y divide-zinc-200 dark:divide-zinc-600 overflow-hidden rounded-lg bg-zinc-100 dark:bg-zinc-700 text-zinc-800 dark:text-white shadow ring-1 ring-zinc-200 dark:ring-zinc-600 mb-8">
            <div class="px-4 py-3 uppercase font-bold">
                Genehmigung
            </div>
            <div class="p-4">
                <div class="grid lg:grid-cols-[1fr_2fr] gap-4">
                    <div class="flex flex-col">
                        <span class="font-bold">AB-OK</span>
                        <span>Vorname Nachname<br />(2025-01-01 01:01:01)</span>
                    </div>
                    <div class="grid lg:grid-cols-[1fr_10rem] gap-4">
                        <div class="flex flex-col">
                            <span class="font-bold">HV-OK</span>
                            <span>Vorname Nachname<br />(2025-01-01 01:01:01)</span>
                        </div>
                        <div class="flex flex-col">
                            <span class="font-bold">Kaution</span>
                            <span>100.00 €</span>
                        </div>
                    </div>
                </div>
                <hr class="my-4 border-zinc-200 dark:border-zinc-600" />
                <div class="grid lg:grid-cols-[1fr_2fr] gap-4">
                    <div class="flex flex-col">
                        <span class="font-bold">Ablehnung</span>
                        <span>Vorname Nachname<br />(2025-01-01 01:01:01)</span>
                    </div>
                    <div class="flex flex-col">
                        <span class="font-bold">Begründung für die Ablehnung</span>
                        <span>Lorem impsum</span>
                    </div>
                </div>
            </div>
        </div>

        <div class="divide-y divide-zinc-200 dark:divide-zinc-600 overflow-hidden rounded-lg bg-zinc-100 dark:bg-zinc-700 text-zinc-800 dark:text-white shadow ring-1 ring-zinc-200 dark:ring-zinc-600 mb-8">
            <div class="px-4 py-3 uppercase font-bold">
                Eckdaten
            </div>
            <div class="p-4 grid gap-4">
                <div class="grid grid-cols-[12rem_1fr] gap-4">
                    <span>Ausleihendes Gremium</span>
                    <span>Lorem ipsum</span>
                </div>
                <div class="grid grid-cols-[12rem_1fr] gap-4">
                    <span>Veranstaltung</span>
                    <span>Lorem ipsum</span>
                </div>
                <div class="grid grid-cols-[12rem_1fr] gap-4">
                    <span>Veranstaltungszeitraum</span>
                    <span>2025-01-01 &ndash; 2025-01-02</span>
                </div>
                <div class="grid grid-cols-[12rem_1fr] gap-4">
                    <span>Projektbeschreibung</span>
                    <span>Lorem ipsum</span>
                </div>
                <div class="grid grid-cols-[12rem_1fr] gap-4">
                    <span>Abholung</span>
                    <span>2025-01-01 01:01:01</span> <span>Ort?</span>
                </div>
                <div class="grid grid-cols-[12rem_1fr] gap-4">
                    <span>Rückgabe</span>
                    <span>2025-01-01 01:01:01</span> <span>Ort?</span>
                </div>
                <div class="grid grid-cols-[12rem_1fr] gap-4">
                    <span>Anmerkungen</span>
                    <span>Lorem ipsum</span>
                </div>
            </div>
        </div>

        <div class="divide-y divide-zinc-200 dark:divide-zinc-600 overflow-hidden rounded-lg bg-zinc-100 dark:bg-zinc-700 text-zinc-800 dark:text-white shadow ring-1 ring-zinc-200 dark:ring-zinc-600">
            <div class="px-4 py-3 uppercase font-bold">
                Gegenstände
            </div>
            <div class="p-4">
                <div>1</div><div>Titel</div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { onMounted, type Ref, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import axios from 'axios'
import SvgIcon from '@jamescoyle/vue-icon'
import {
  mdiCartRemove
} from '@mdi/js';

const { t } = useI18n()


</script>

<style scoped>
tr {
    display: grid;
    grid-template-columns: 12rem 1fr;
    gap: .5rem;
    border-bottom: 1px solid #52525b;
}

tr:first-of-type {
    border-top: 1px solid #52525b;
}

th {
    text-align: left;
}

th, td {
    padding: .25rem;
}

#submitted {
    transform: translate(2,2);
}

#withdrawn {
    transform: translate(2,102);
}
</style>