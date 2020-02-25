async function initMongo() {
    console.log('Initialising MongoDB...')
    let success = false
    while (!success) {
      try {
        client = await MongoClient.connect(mongoURL, { useNewUrlParser: true })
        success = true
      } catch {
        console.log('Error connecting to MongoDB, retrying in 1 second')
        await new Promise(resolve => setTimeout(resolve, 1000))
      }
    }
    console.log('MongoDB initialised')
    return client.db(client.s.options.dbName).collection('notes')
  }