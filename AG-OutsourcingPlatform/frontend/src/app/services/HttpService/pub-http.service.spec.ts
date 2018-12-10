import { TestBed, inject } from '@angular/core/testing';

import { PubHttpService } from './pub-http.service';

describe('PubHttpService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [PubHttpService]
    });
  });

  it('should be created', inject([PubHttpService], (service: PubHttpService) => {
    expect(service).toBeTruthy();
  }));
});
