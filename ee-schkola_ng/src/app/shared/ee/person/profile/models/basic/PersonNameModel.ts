import {PersonName} from '../../../../schkola/person/PersonApiBase';

export class PersonNameModel {

    public name: PersonName = new PersonName();

    public firstname = ['firstname', 'string', typeof this.name.first];
    public lastname = ['lastname', 'string', typeof this.name.last];
}
