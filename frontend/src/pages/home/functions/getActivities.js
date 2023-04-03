/**
 Sends a request to the Bumbal API to retrieve a list of activities.
 @returns {Promise<Response>} - The response from the API containing a list of activities.
 */
 async function getActivities(offset) {
    const token = localStorage.getItem('token')
    console.log('token ' + token)
  
    const requestOptions = {
      method: 'PUT',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`, // Add token to Bearer Authorization when sending GET signOut request.
      },
      body: JSON.stringify({
        "options": {
            "include_address": true,
            "include_depot_address": true
        },
          "sorting_column": "id",
          "as_list": true,
          "count_only": false,
          "limit": 100,
          "offset": offset
      
      })
    };
  
    const response = await fetch('https://sep202302.bumbal.eu/api/v2/activity', requestOptions);
  
  
    return response;
}

export default getActivities
