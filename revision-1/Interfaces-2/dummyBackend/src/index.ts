import express from 'express';

const app = express();

app.get('/sendit', (req, res) => {
   res.json({
      message: 'Hello there, here is your response',
   });
});

app.listen(3000, () => console.log('Listening on Port 3000'));
