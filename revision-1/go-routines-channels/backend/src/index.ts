import express from 'express';

const app = express();

app.get('/home', async (req, res) => {
   await new Promise((resolve) => setTimeout(resolve, 1000));
   return res.status(200).json({
      message: 'Captured',
   });
});

app.listen(3000, () => console.log('Listening on PORT 3000'));
