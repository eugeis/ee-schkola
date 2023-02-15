import { DatePipe } from '@angular/common';
import { Pipe, PipeTransform } from '@angular/core';
import { TranslateService } from '@ngx-translate/core';
import {TableDataService} from '@template/services/data.service';

@Pipe({
    name: 'DateTimeTranslation',
    pure: false
})
export class DateTimeTranslationPipe implements PipeTransform {

    constructor(private translateService: TranslateService) {}

    transform(value: Date, pattern: string = localStorage.getItem('pattern') ? localStorage.getItem('pattern') : 'shortDate'): any {
        const datePipe: DatePipe = new DatePipe(
            this.translateService.currentLang ? this.translateService.currentLang : this.translateService.getDefaultLang());
        return datePipe.transform(value, pattern);
    }
}
