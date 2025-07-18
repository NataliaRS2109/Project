<script setup lang="ts">
    import type { Item } from '../models/item.model'; // Importa el tipo Item desde el modelo
    import { ref, computed } from 'vue'
    // Define una interfaz Props para tipar las propiedades del componente.
    interface Props { // Define una interfaz Props para tipar las propiedades del componente.
        item: Item[];
    }

    const sortColumn = ref<string | null>(null);
    const sortDirection = ref<'asc' | 'desc' | null>(null);

    const sortedItems = computed(() => {
        if (!sortColumn.value) {
        return props.item;
        }

        const sorted = [...props.item].sort((a, b) => {
        const valueA = a[sortColumn.value as keyof Item];
        const valueB = b[sortColumn.value as keyof Item];

        if (valueA < valueB) {
            return sortDirection.value === 'asc' ? -1 : 1;
        }
        if (valueA > valueB) {
            return sortDirection.value === 'asc' ? 1 : -1;
        }
        return 0;
        });
        return sorted;
    });

    const sort = (column: string) => {
        if (sortColumn.value === column) {
        sortDirection.value = sortDirection.value === 'asc' ? 'desc' : 'asc';
        } else {
        sortColumn.value = column;
        sortDirection.value = 'asc';
        }
    }; 

    const props = defineProps<Props>(); // Define las propiedades del componente usando la interfaz Props. Esto permite que el componente reciba un array de objetos Country como propiedad.
</script>

<template>
    <div class="relative flex flex-col w-full h-full overflow-auto bg-white shadow-md rounded-lg bg-clip-border mt-4">
        <table class="w-full table-auto min-w-max text-left">
            <thead>
                <tr>
                    <th class="p-4 cursor-pointer border-b border-slate-300 bg-gray-300" @click="sort('ticker')">
                        <p class="text-sm font-normal leading-none text-slate-800">
                            TICKER
                            <span v-if="sortColumn === 'ticker'">
                                {{ sortDirection === 'asc' ? '⬆️' : '⬇️' }} <!-- Agrega una flecha de acuerdo a como se esté ordenando la columna -->
                            </span>
                        </p>
                    </th>            
                    <th class="p-4 cursor-pointer border-b border-slate-300 bg-gray-300" @click="sort('company')">
                        <p class="text-sm font-normal leading-none text-slate-800">
                            COMPANY 
                            <span v-if="sortColumn === 'company'">
                                {{ sortDirection === 'asc' ? '⬆️' : '⬇️' }}
                            </span>
                        </p>
                    </th>
                    <th class="p-4 cursor-pointer border-b border-slate-300 bg-gray-300" @click="sort('brokerage')">
                        <p class="text-sm font-normal leading-none text-slate-800">
                            BROKERAGE
                            <span v-if="sortColumn === 'brokerage'">
                                {{ sortDirection === 'asc' ? '⬆️' : '⬇️' }}
                            </span>
                        </p>
                    </th>
                    <th class="p-4 cursor-pointer border-b border-slate-300 bg-gray-300" @click="sort('action')">
                        <p class="text-sm font-normal leading-none text-slate-800">
                            ACTION
                            <span v-if="sortColumn === 'action'">
                                {{ sortDirection === 'asc' ? '⬆️' : '⬇️' }}
                            </span>
                        </p>
                    </th>
                    <th class="p-4 cursor-pointer border-b border-slate-300 bg-gray-300" @click="sort('rating_from')">
                        <p class="text-sm font-normal leading-none text-slate-800">
                            RATING_FROM
                            <span v-if="sortColumn === 'rating_from'">
                                {{ sortDirection === 'asc' ? '⬆️' : '⬇️' }}
                            </span>
                        </p>
                    </th>
                    <th class="p-4 cursor-pointer border-b border-slate-300 bg-gray-300" @click="sort('rating_to')">
                        <p class="text-sm font-normal leading-none text-slate-800">
                            RATING_TO
                            <span v-if="sortColumn === 'rating_to'">
                                {{ sortDirection === 'asc' ? '⬆️' : '⬇️' }}
                            </span>
                        </p>
                    </th>
                    <th class="p-4 cursor-pointer border-b border-slate-300 bg-gray-300" @click="sort('target_from')">
                        <p class="text-sm font-normal leading-none text-slate-800">
                            TARGET_FROM
                            <span v-if="sortColumn === 'target_from'">
                                {{ sortDirection === 'asc' ? '⬆️' : '⬇️' }}
                            </span>
                        </p>
                    </th>
                    <th class="p-4 cursor-pointer border-b border-slate-300 bg-gray-300" @click="sort('target_to')">
                        <p class="text-sm font-normal leading-none text-slate-800">
                            TARGET_TO
                            <span v-if="sortColumn === 'target_to'">
                                {{ sortDirection === 'asc' ? '⬆️' : '⬇️' }}
                            </span>
                        </p>
                    </th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="item in sortedItems" class="hover:bg-gray-100">
                    <td class="p-4 border-b border-slate-200">
                        <p class="block text-sm text-slate-800">
                            {{ item.ticker }}
                        </p>
                    </td>
                    <td class="p-4 border-b border-slate-200">
                        <p class="block text-sm text-slate-800">
                            {{ item.company }}
                        </p>
                    </td>
                    <td class="p-4 border-b border-slate-200">
                        <p class="block text-sm text-slate-800">
                            {{ item.brokerage }}
                        </p>
                    </td>
                    <td class="p-4 border-b border-slate-200">
                        <p class="block text-sm text-slate-800">
                            {{ item.action }}
                        </p>
                    </td>
                    <td class="p-4 border-b border-slate-200">
                        <p class="block text-sm text-slate-800">
                            {{ item.rating_from }}
                        </p>
                    </td>
                    <td class="p-4 border-b border-slate-200">
                        <p class="block text-sm text-slate-800">
                            {{ item.rating_to }}
                        </p>
                    </td>
                    <td class="p-4 border-b border-slate-200">
                        <p class="block text-sm text-slate-800">
                            {{ item.target_from }}
                        </p>
                    </td>
                    <td class="p-4 border-b border-slate-200">
                        <p class="block text-sm text-slate-800">
                            {{ item.target_to }}
                        </p>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>


