import { TestBed } from '@angular/core/testing';

import { UpdateStatusService } from './update-status.service';

describe('UpdateStatusService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: UpdateStatusService = TestBed.get(UpdateStatusService);
    expect(service).toBeTruthy();
  });
});
