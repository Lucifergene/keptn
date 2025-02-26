import { TestBed } from '@angular/core/testing';
import { AppModule } from '../app.module';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { FeatureFlagsService } from './feature-flags.service';
import { ApiService } from './api.service';
import { ApiServiceMock } from './api.service.mock';
import { DataService } from './data.service';
import { take } from 'rxjs/operators';

describe('FeatureFlagsService', () => {
  let service: FeatureFlagsService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [],
      imports: [AppModule, HttpClientTestingModule],
      providers: [{ provide: ApiService, useClass: ApiServiceMock }],
    });

    service = TestBed.inject(FeatureFlagsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should set feature flags', async () => {
    const dataService = TestBed.inject(DataService);
    dataService.loadKeptnInfo();
    const flags = await service.featureFlags$.pipe(take(1)).toPromise();
    expect(flags).toEqual({
      RESOURCE_SERVICE_ENABLED: false,
    });
  });
});
