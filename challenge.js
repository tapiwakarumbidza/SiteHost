const https = require('https');

const API_BASE = 'https://api.recruitment.shq.nz';
const API_KEY = 'h523hDtETbkJ3nSJL323hjYLXbCyDaRZ';
const CLIENT_ID = '100';

function makeRequest(path) {
  return new Promise((resolve, reject) => {
    const url = `${API_BASE}${path}?api_key=${API_KEY}`;
    
    https.get(url, (res) => {
      let data = '';
      
      res.on('data', (chunk) => {
        data += chunk;
      });
      
      res.on('end', () => {
        try {
          resolve(JSON.parse(data));
        } catch (e) {
          reject(e);
        }
      });
    }).on('error', reject);
  });
}

async function main() {
  try {
    console.log('SiteHost Recruitment Challenge - Part One\n');
    console.log('Fetching domains for Business Systems International (Client ID: ' + CLIENT_ID + ')...\n');
    
    // Part One: Get domains for the client
    const domainsResponse = await makeRequest(`/domains/${CLIENT_ID}`);
    console.log('Domains:');
    console.log(JSON.stringify(domainsResponse, null, 2));
    console.log('\n');
    
    // Get DNS records for each zone
    const domains = Array.isArray(domainsResponse) ? domainsResponse : domainsResponse.domains;
    if (domains && domains.length > 0) {
      for (const domain of domains) {
        console.log(`Domain: ${domain.name}`);
        
        if (domain.zones && domain.zones.length > 0) {
          for (const zone of domain.zones) {
            console.log(`  Zone: ${zone.name}`);
            
            const zoneId = zone.uri.split('/').pop();
            const recordsResponse = await makeRequest(`/zones/${zoneId}`);
            console.log(`  DNS Records:`);
            console.log(JSON.stringify(recordsResponse, null, 2));
            console.log('');
          }
        } else {
          console.log('  No zones configured');
        }
        console.log('');
      }
    }
    
  } catch (error) {
    console.error('Error:', error.message);
    process.exit(1);
  }
}

main();
