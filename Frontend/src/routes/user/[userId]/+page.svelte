<script>
    import { page } from '$app/stores';
    import {onMount} from "svelte";

    $: userId = $page.params.userId;
    let user;

    const getUserByID = async (id) => {
        const res = fetch(`http://localhost:8080/api/user/:${id}`);
        user = await res.json();
    }

    onMount(getUserByID);
</script>

<h1>User ID: {userId}</h1>
