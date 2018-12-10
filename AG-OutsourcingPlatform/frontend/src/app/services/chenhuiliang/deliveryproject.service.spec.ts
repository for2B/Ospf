import { TestBed, inject } from '@angular/core/testing';

import { DeliveryprojectService } from './deliveryproject.service';

describe('DeliveryprojectService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [DeliveryprojectService]
    });
  });

  it('should be created', inject([DeliveryprojectService], (service: DeliveryprojectService) => {
    expect(service).toBeTruthy();
  }));
});
