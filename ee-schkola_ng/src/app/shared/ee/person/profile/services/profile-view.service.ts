import {Injectable} from '@angular/core';
import {TemplateService} from '../../../template/services/template.service';
import {PersonNameModel} from '../models/basic/PersonNameModel';
import {GenderEnumModel} from '../models/enums/GenderEnumModel';

@Injectable()
export class ProfileViewService extends TemplateService {

    personName: PersonNameModel = new PersonNameModel();

    elementNameWithValue = [['gender', 'enum', GenderEnumModel], this.personName.firstname, this.personName.lastname,
        ['birthName', 'string'], ['birthday', 'datetime']];

    pageElement = ['Person'];

    tabElement = ['Profile', 'Church'];

    pageName = 'ProfileComponent';

    formArrayName = ['Actions'];
}
