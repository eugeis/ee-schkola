import {Injectable} from '@angular/core';
import {FormService} from '../../../template/services/form.service';
import {PersonNameModel} from '../../models/basic/PersonNameModel';
import {GenderEnumModel} from '../../models/enums/GenderEnumModel';
import {AddressModel} from '../../models/basic/AddressModel';

@Injectable()
export class ProfileViewService extends FormService {

    personName: PersonNameModel = new PersonNameModel();

    address: AddressModel = new AddressModel();

    form = this.fb.group({
        gender: 'UNKNOWN',
        firstname: '',
        lastname: '',
        birthName: '',
        birthday: this.dateNow,
        street: '',
        suite: '',
        code: '',
        city: '',
        country: '',
    });

    formArrayType = ['enum', 'string', 'string', 'string', 'datetime', 'string', 'string', 'string', 'string', 'string'];

    formArrayName = ['Actions', 'gender', 'firstname', 'lastname', 'birthName', 'birthday', 'street', 'suite', 'code', 'city', 'country'];

    formEnumGender = this.loadEnumElement(GenderEnumModel);

    pageElement = ['Person'];

    tabElement = ['Profile', 'Church'];

    pageName = 'ProfileComponent';
}
