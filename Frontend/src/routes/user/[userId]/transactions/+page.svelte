<script>
    import {page} from "$app/stores";
    import {onMount} from "svelte";

    $: userId = $page.params.userId;

    let user;
    let transactions = [];

    const getUserByID = async () => {
        const res = fetch(`http://localhost:3030/api/user/${userId}`, { method: 'GET' });
        user = (await res).json();
    }

    const getAllTransactions = async () => {
        const res = fetch(`http:/localhost:3030/api/transactions`, { method: 'GET' });
        transactions = (await res).json();
    }

    const getData = async () => {
        await getUserByID();
        await getAllTransactions();
    }
    onMount(getData);
</script>

<div class="transactions">
    {#each transactions as {id, amount, category, type, created_at}}
        <div class="transaction">
            <h1>Transaction: {id}</h1>
            <div>
                Amount: {amount},
                Category: {category},
                Type: {type},
                Created at: {created_at}.
            </div>
        </div>
    {/each}
</div>

<style>
    .transactions {
        display: flex;
        overflow-x: auto;
    }

    .transaction {
        margin: auto;
    }
</style>
