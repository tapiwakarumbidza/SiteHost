const https = require('https');
const http = require('http');

const IP_ADDRESS = '120.138.30.179';
const HOST = 'site.recruitment.shq.nz';

function makeRequest(host, ip) {
  return new Promise((resolve, reject) => {
    const options = {
      hostname: ip,
      port: 80,
      path: '/',
      method: 'GET',
      headers: {
        'Host': host
      },
      timeout: 5000
    };

    const req = http.request(options, (res) => {
      let data = '';

      res.on('data', (chunk) => {
        data += chunk;
      });

      res.on('end', () => {
        resolve(data);
      });
    });

    req.on('error', (error) => {
      reject(error);
    });

    req.on('timeout', () => {
      req.destroy();
      reject(new Error('Request timeout'));
    });

    req.end();
  });
}

async function main() {
  try {
    console.log(`Attempting to access ${HOST} using IP ${IP_ADDRESS}...\n`);
    
    const html = await makeRequest(HOST, IP_ADDRESS);
    
    console.log('Success! Website content retrieved:\n');
    console.log(html);
    
    // Extract the hidden code from HTML comments in the head
    const headMatch = html.match(/<head[^>]*>([\s\S]*?)<\/head>/i);
    if (headMatch) {
      const headContent = headMatch[1];
      const codeMatch = headContent.match(/<!--\s*This is what you're looking for:\s*([A-Z0-9_-]+)\s*-->/);
      if (codeMatch) {
        console.log('\n\n=== FOUND HIDDEN CODE ===');
        console.log(`Code: ${codeMatch[1]}`);
      }
    }
    
  } catch (error) {
    console.error('Error:', error.message);
    process.exit(1);
  }
}

main();
