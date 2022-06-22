import {Address} from '../../../schkola/person/PersonApiBase';

export class AddressModel {

    public address: Address = new Address();

    public street = ['street', 'string', typeof this.address.street];
    public suite = ['suite', 'string', typeof this.address.suite];
    public city = ['city', 'string', typeof this.address.city];
    public code = ['code', 'string', typeof this.address.code];
    public country = ['country', 'string', typeof this.address.country];
}
