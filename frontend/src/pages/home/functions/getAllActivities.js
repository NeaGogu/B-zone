import getActivities from "./getActivities"

async function getAllActivities() {
    const token = localStorage.getItem('token')

    // first see how many activities there are
    let  activities =  await fetch('https://sep202302.bumbal.eu/api/v2/activity', {
      method: 'PUT',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`, // Add token to Bearer Authorization when sending GET signOut request.
      },
      body: JSON.stringify({
        "count_only": true
      })
    }).then((response)=>{return response.json()})
  
    let datatop = []
    // for loop and get activities based on offset
    for (let i = 0; i < activities.count_filtered;i = i + 100) {
        const response = await getActivities(i);
  
      // If token is invalid, take the user to log in page.
      if (response.status === 401) {
        alert('Expired or Invalid Token')
        localStorage.removeItem('token')
        window.location.reload()
      }
  
      const data = await response.json();
      datatop.push(...data.items)
    }

    return datatop
}

export default getAllActivities