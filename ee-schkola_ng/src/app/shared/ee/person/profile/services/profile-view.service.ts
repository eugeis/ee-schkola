import {Injectable} from '@angular/core';
import {TemplateService} from '../../../template/services/template.service';
import {PersonNameModel} from '../../models/basic/PersonNameModel';
import {GenderEnumModel} from '../../models/enums/GenderEnumModel';
import {AddressModel} from '../../models/basic/AddressModel';

@Injectable()
export class ProfileViewService extends TemplateService {

    personName: PersonNameModel = new PersonNameModel();

    address: AddressModel = new AddressModel();

    elementNameWithValue = [['gender', 'enum', GenderEnumModel], this.personName.firstname, this.personName.lastname,
        ['birthName', 'string'], ['birthday', 'datetime'],
        this.address.street, this.address.suite, this.address.code, this.address.city, this.address.country];

    pageElement = ['Person'];

    tabElement = ['Profile', 'Church'];

    pageName = 'ProfileComponent';

    formArrayName = ['Actions'];
}
