describe('infra policy groups', () => {
  const now = Cypress.moment().format('MMDDYYhhmmss');
  const cypressPrefix = 'infra';
  let adminIdToken = '';
  const serverID = 'chef-server-dev-test';
  const serverName = 'chef server dev';
  const orgID = 'chef-org-dev';
  const orgName = '4thcoffee';
  const serverFQDN = 'ec2-34-219-25-251.us-west-2.compute.amazonaws.com';
  const serverIP = '34.219.25.251';
  const adminUser = 'chefadmin';
  const adminKey = Cypress.env('AUTOMATE_INFRA_ADMIN_KEY').replace(/\\n/g, '\n');
  let policies: any;

  before(() => {
    cy.adminLogin('/').then(() => {
      const admin = JSON.parse(<string>localStorage.getItem('chef-automate-user'));
      adminIdToken = admin.id_token;

      cy.request({
        auth: { bearer: adminIdToken },
        failOnStatusCode: false,
        method: 'POST',
        url: '/api/v0/infra/servers',
        body: {
          id: serverID,
          name: serverName,
          fqdn: serverFQDN,
          ip_address: serverIP
        }
      }).then((resp) => {
        if (resp.status === 200 && resp.body.ok === true) {
          return;
        } else {
          cy.request({
            auth: { bearer: adminIdToken },
            method: 'GET',
            url: `/api/v0/infra/servers/${serverID}`,
            body: {
              id: serverID
            }
          });
        }
      });

      cy.request({
        auth: { bearer: adminIdToken },
        failOnStatusCode: false,
        method: 'POST',
        url: `/api/v0/infra/servers/${serverID}/orgs`,
        body: {
          id: orgID,
          server_id: serverID,
          name: orgName,
          admin_user: adminUser,
          admin_key: adminKey
        }
      }).then((response) => {
        if (response.status === 200 && response.body.ok === true) {
          return;
        } else {
          cy.request({
            auth: { bearer: adminIdToken },
            method: 'GET',
            url: `/api/v0/infra/servers/${serverID}/orgs/${orgID}`,
            body: {
              id: orgID,
              server_id: serverID
            }
          });
        }
      });
      cy.visit(`/infrastructure/chef-servers/${serverID}/organizations/${orgID}`);
      cy.get('app-welcome-modal').invoke('hide');
    });
    cy.restoreStorage();
  });

  beforeEach(() => {
    cy.restoreStorage();
  });

  afterEach(() => {
    cy.saveStorage();
  });

  function getPolicyGroups() {
    return cy.request({
      auth: { bearer: adminIdToken },
      failOnStatusCode: false,
      method: 'GET',
      url: `/api/v0/infra/servers/${serverID}/orgs/${orgID}/policyfiles`
    });
  }

  function checkResponse(response: any) {
    if (response.body.policies.length === 0) {
      cy.get('[data-cy=empty-list]').should('be.visible');
    } else {
      policies = response.body.policies;
      cy.get('[data-cy=policy-group-table-container] chef-th').contains('Policy Group');
      cy.get('[data-cy=policy-group-table-container] chef-th').contains('Number of Policyfiles');
      return true;
    }
  }

  describe('infra policy group list page', () => {
    it('displays org details', () => {
      cy.get('.page-title').contains(orgName);
    });

    // policy group tabs specs
    it('can switch to policy group tab', () => {
      cy.get('.nav-tab').contains('Policy Groups').click();
    });

    it('can check if policy group has list or not', () => {
      getPolicyGroups().then((response) => {
        checkResponse(response);
      });
    });
  });
});