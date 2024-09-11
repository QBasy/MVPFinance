<script>
    import { page } from '$app/stores';
    import {onMount} from "svelte";

    $: userId = $page.params.userId;
    let user;

    const getUserByID = async () => {
        const res = fetch(`http://localhost:3030/api/user/:${userId}`);
        user = (await res).json();
    }

    onMount(getUserByID);
</script>

<h1>User ID: {userId}</h1>
