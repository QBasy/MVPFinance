<script>
    import { page } from '$app/stores';
    import {onMount} from "svelte";

    $: userId = $page.params.userId;
    $: transactionId = $page.params.transactionId;

    let user;
    let transaction;

    const getUserByID = async () => {
        const res = fetch(`http://localhost:3030/api/user/${userId}`);
        user = (await res).json();
    }

    const getTransactionById = async () => {
        const res = fetch(`http://localhost:8080/api/transaction/${transactionId}`);
        transaction = (await res).json();
    }

    async function getData() {
        await getTransactionById()
        await getUserByID()
    }

    onMount(getData)
</script>

<div><h1>User ID: {userId}</h1></div>

<div>
    <h2>Transaction ID: {transactionId}</h2>
</div>
