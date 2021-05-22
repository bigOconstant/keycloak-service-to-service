#include <iostream>
#include <cpr/cpr.h>
#include <nlohmann/json.hpp>
using std::cout;
using std::endl;

int main() {
    cpr::Response r = cpr::Post(cpr::Url{std::getenv("TOKENURL")},
                   cpr::Payload{{"grant_type", "client_credentials"},{"client_id",std::getenv("CLIENTID")},{"client_secret", std::getenv("CLIENTSECRET")}},
                   cpr::Header{{"Content-Type", "application/x-www-form-urlencoded"}});

   if(r.error) {
       cout<<"error making request"<<endl;
       exit (EXIT_FAILURE);
   }

   nlohmann::json token = nlohmann::json::parse(r.text);

   cout<<"Token expires in:"<<token["expires_in"]<< " seconds \n"<<endl;
   cout<<"Token:"<<token["access_token"]<<endl;

    return 0;
}